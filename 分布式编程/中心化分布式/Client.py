import time #时间
import queue #队列

from    multiprocessing.managers import BaseManager #分布式管理器




class QueueMananger(BaseManager):#创建类继承式分布式管理器
    pass

QueueMananger.register("get_task_queue")#任务队列
QueueMananger.register("get_result_queue")#结果队列

print("客户端开始")
client=QueueMananger(address=("127.0.0.1",55000),authkey=b"123456")
client.connect()#链接服务器
for i in range(10):
    n=client.get_task_queue().get()
    print("客户端取出数据",n)
    time.sleep(1)
    client.get_result_queue().put(n*n)

print("客户端结束")