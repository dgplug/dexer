#!/usr/bin/python3

import os, random, sys

number_of_files = random.randint(100,200)
number_of_folders = random.randint(5,15)
working_dir = os.path.join(os.getcwd(),sys.argv[1])

for i in range(number_of_folders):
    changed_dir = os.path.join(working_dir, 'folder' + str(i))
    os.mkdir(changed_dir)
    for j in range(number_of_files):
        file_name = os.path.join(changed_dir, 'file' + str(j));
        with open(file_name, 'wb') as fout:
            fout.write(os.urandom(1024))
    number_of_files = random.randint(100,200)


