# cat post_gen_project.py
import os
import shutil

print(os.getcwd())  # prints /absolute/path/to/{{cookiecutter.lambda_name}}


def remove(filepath):
    if os.path.isfile(filepath):
        os.remove(filepath)
    elif os.path.isdir(filepath):
        shutil.rmtree(filepath)


use_dynamodb = '{{cookiecutter.use_dynamodb}}' == 'y'

if not use_dynamodb:
    build_path = os.path.join(os.getcwd(), 'build')
    remove(build_path)
    dynamodb_file = os.path.join(os.getcwd(), 'pkg', 'storage', 'dynamodb.go')
    remove(dynamodb_file)
