import datetime #时间日期类
import hashlib  #信息安全加密解密


class DaDaMessage:
    def __init__(self,data):#初始化
        self.hash=None#自身的哈希
        self.prev_hash=None #上一个信息记录的哈希
        self.timestamp=datetime.datetime.now()#交易时间
        self.data=data #交易信息
        self.payload_hash=self.__hash__payload()#交易后的哈希

    def __hash__payload(self):#对于交易时间与交易数据进行哈希计算
        return hashlib.sha512(   (str(self.timestamp )+str(self.data)).encode("utf-8")   ).hexdigest()#取得数据的哈希

    def __hash__message(self):#对于交易进行锁定
        return hashlib.sha512((str(self.prev_hash) + str(self.payload_hash)).encode("utf-8")).hexdigest()  # 取得数据的哈希

    def seal(self):#密封
        self.hash=self.__hash__message() #对应数据锁定，对于交易前的锁定

    def validate(self):#验证
        if self.payload_hash !=self.__hash__payload():#判断是否有人修改
           raise  InvalidMessage("交易数据与时间被修改"+str(self))
        if self.hash!=self.__hash__message():#判断消息链
            raise InvalidMessage("交易的哈希链被修改"+str(self))


    def __repr__(self): #返回对象的基本信息
        mystr= "hash:{},prev_hash:{},data:{} ".format(self.hash,self.prev_hash,self.data)
        return mystr

    def link(self,Message):
        self.prev_hash=Message.hash #链接

class InvalidMessage(Exception):#异常类型
    def __init__(self,*args,**kwargs):#参数
        Exception.__init__(self,*args,**kwargs)

if  __name__=="__main__":#单独模块测试
    try:
        m1= DaDaMessage("zhougl pay 51cto 10 coins")
        m2= DaDaMessage("zhougl pay debao 20 coins")
        m3= DaDaMessage("zhougl pay jikexuyuan 30 coins")#交易记录
        m4= DaDaMessage("zhougl pay jinwei 40 coins")#交易记录


        m1.seal()
        m2.link(m1)#链接
        m2.seal()
        m3.link(m2)#链接
        m3.seal() #交易记录密封
        m4.link(m3)#链接
        m4.seal() #交易记录密封

        #修改数据，篡改数据
        m2.data="你妹的直播平台"
        m2.prev_hash="哈哈哈哈哈，你哥哥"



        print (m1)
        print (m2)
        print (m3)#显示信息
        print (m4)#显示信息

        m1.validate()
        m2.validate()
        m3.validate()#校验
    except InvalidMessage as e:
        print(e)