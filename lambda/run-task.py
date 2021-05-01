import json
import os
import boto3

ecs = boto3.client('ecs')

# Lambdaの[設定]→[環境変数]で設定する
ECS_CLUSTER = os.environ["ECS_CLUSTER"]
TASK_DEFINITION = os.environ["TASK_DEFINITION"]
# 起動時のネットワークのサブネット
SUBNET_ID_1 = os.environ["SUBNET_ID_1"]
SUBNET_ID_2 = os.environ["SUBNET_ID_2"]


# イベントトリガーで起動する関数
def lambda_handler(event, context):
    # ECSタスクの実行
    response = ecs.run_task(
        cluster=ECS_CLUSTER,
        taskDefinition=TASK_DEFINITION,
        launchType='FARGATE',
        networkConfiguration={
            'awsvpcConfiguration': {
                'subnets': [
                    SUBNET_ID_1,
                    SUBNET_ID_2
                ],
                'assignPublicIp': 'ENABLED',
            }
        },
    )

    # エラーの検知
    print(response)
    failures = response['failures']
    if len(failures) != 0:
        print(failures)
        return {
            'statusCode': 500,
            'body': json.dumps('error occurred at RunTask')
        }

    # 正常終了
    return {
        'statusCode': 200,
        'body': json.dumps('successfull RunTask')
    }
