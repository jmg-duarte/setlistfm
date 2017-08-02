import sys
import re

with open(sys.argv[1]) as result:
    pattern = re.compile(r'\s+')
    sentence = re.sub(pattern, '', result.read())
    with open(sys.argv[1]+"_clean", "w+") as out:
        out.write(sentence)