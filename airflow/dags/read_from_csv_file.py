import os
import csv
import json
import random

from airflow.constants import constants

from airflow.models import DAG
from airflow.utils.dates import days_ago
from airflow.contrib.hooks.fs_hook import FSHook
from airflow.contrib.sensors.file_sensor import FileSensor
from airflow.operators.python_operator import PythonOperator


args = {
    'owner':  constants.OWNER,
    'start_date': days_ago(1)
}

dag = DAG(
    default_args=args,
    dag_id='my_sample-dag',
    schedule_interval=constants.DAG_RUN_TIME_INTRVAL
)


def read_file(**context):
    hook = FSHook('fs_default')
    base_path = hook.get_path()
    file_path = os.path.join(base_path, constants.CSV_FILE_NAME)

    json_data = {}
    with open(file_path, 'r') as file:
        reader = csv.DictReader(file)

        for index, item in enumerate(reader, start=0):   # default is zero
            json_data[index] = item

    with open(base_path + constants.JSON_FILE_NAME, 'w') as json_file:
        json.dump(json_data, json_file)

    print(json_data)


with dag:
    sensing_task = FileSensor(
        poke_interval=10,
        filepath='laptop.csv',
        task_id='check_file_existance',
        fs_conn_id='fs_default'
    )

    read_file = PythonOperator(
        task_id='read_file',
        python_callable=read_file,
        provide_context=True
    )

    sensing_task >> read_file
