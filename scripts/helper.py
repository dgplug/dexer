#!/usr/bin/python3

import os, random, sys
from faker import Faker

number_of_files = random.randint(5,10)
number_of_folders = random.randint(1,5)
working_dir = os.path.join(os.getcwd(),sys.argv[1])
fake = Faker()

for i in range(number_of_folders):
    changed_dir = os.path.join(working_dir, 'folder' + str(i))
    os.mkdir(changed_dir)
    for j in range(number_of_files):
        file_name = os.path.join(changed_dir, 'file' + str(j))
        with open(file_name, 'w') as fout:
            fout.writelines(fake.text())
    number_of_files = random.randint(5,10)


