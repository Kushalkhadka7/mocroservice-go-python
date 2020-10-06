import os
import csv
import json

from airflow.constants import constants

from airflow.contrib.hooks.fs_hook import FSHook


class FileOperation:
    """
    Handle all the operations related to the file.
    Read the data from the csv file process it.
    Write the processed data to the json file.
    Read the data available in the json file.
    """

    def __init__(self, base_path) -> str:
        self.base_path = base_path

    def read_from_json_file(self, file_name):
        file_path = os.path.join(self.base_path, file_name)

        with open(file_path, 'r') as json_file:
            loaded_file = json.load(json_file)
        return loaded_file

    def read_csv_write_json(self, dest_file, source_file):
        hook = FSHook('fs_default')
        base_path = hook.get_path()
        file_path = os.path.join(base_path, source_file)

        json_data = {}
        with open(file_path, 'r') as file:
            reader = csv.DictReader(file)

            for index, item in enumerate(reader, start=0):   # default is zero
                json_data[index] = item

        with open(base_path + dest_file, 'w') as json_file:
            json.dump(json_data, json_file)

        return json_data
