# Adapted from https://stackoverflow.com/questions/33365621:
import re

# Read in examples.fs:
with open('examples.fs', 'r') as fs:
    examples = fs.read()

with open('README.md', 'r') as md:
        readme = md.read()
        ntext = re.sub(r'(?<=BEGIN EXAMPLES \-\-\>\n)(.*)(?=\<\!\-\- END EXAMPLES)',
                       examples,
                       readme,
                       flags=re.M|re.DOTALL)

with open('README.md', 'w') as fout:
    fout.write(ntext)
