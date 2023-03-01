# Adapted from https://stackoverflow.com/questions/33365621:
import re

def execute(code):
    import subprocess
    proc = subprocess.Popen(code,
                            shell=True,
                            stdout=subprocess.PIPE,
                            stderr=subprocess.PIPE)
    out, _ = proc.communicate()  # Optimism!!!
    return out.decode('utf-8')

steps = ["go build ."]

steps_txt = ("\n" +
             "".join([f"$ {s}\n{execute(s)}\n" for s in steps]) +
             "\n")

with open('README.md', 'r') as md:
        readme = md.read()
        examples = ("\n    " +
                    "\n    ".join([l.rstrip() for l in steps_txt.split('\n')]) +
                    "\n")
        ntext = re.sub(r'(?<=BEGIN EXAMPLES \-\-\>\n)(.*)(?=\<\!\-\- END EXAMPLES)',
                       examples,
                       readme,
                       flags=re.M|re.DOTALL)

with open('README.md', 'w') as fout:
    fout.write(ntext)
