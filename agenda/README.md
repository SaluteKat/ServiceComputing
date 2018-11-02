# 服务计算——goAgenda

---

## 一、安装使用cobra  
这里是参考老师的指令:  
- 首先我们需要进入 ```$GOPATH/src/golang.org/x```文件夹中，使用命令：  
    ```
    git clone https://github.com/golang/sys.git
    git clone https://github.com/golang/text.git
    ```  
    当这两条指令克隆完成之后再使用：  
    ```
    go install github.com/spf13/cobra/cobra
    ```   
    这样就完成了cobra的安装  

- 使用cobra创建命令：  
    先使用```cobra init```完成初始化，如果需要注册register命令可以使用：  
    ```
    cobra add register
    ```
    之后我们可以在cmd中完成命令的创建和行为的补充，在entity中实现命令逻辑方面的东西。  

--- 

## 二、完成的Agenda命令：    
我们团队完成了如下命令： 

- agenda help
- agenda register --username=User --password=****** --email=xxx@xx.com  
- agenda login --username=User --password=******  
- agenda logout  
- agenda createMeeting --title=Title --participants=Part --startTime=0000-00-00/00:00 --endTime=0000-00-00/00:00  
- agenda deleteMeeting --title=Title  

---  
## 三、命令测试：  

- agenda help   
![image1](https://img-blog.csdnimg.cn/20181102131556878.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0VtaWx5Qmx1c2U=,size_16,color_FFFFFF,t_70)  

- agenda register  
    这里我先注册了很多个用户：  
    ![image2](https://img-blog.csdnimg.cn/2018110213190696.png)   
    这是注册用户的信息。注册之后的用户记录在```data/User.json```中：  
    ![image3](https://img-blog.csdnimg.cn/20181102132011891.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0VtaWx5Qmx1c2U=,size_16,color_FFFFFF,t_70)   

- agenda login  
    这里测试登陆用户名为haha的用户：  
    如果用户名不存在，则出现报错：  
    ![image4](https://img-blog.csdnimg.cn/20181102132201172.png)
    如果用户名正确，密码错误，也会出现报错：  
    ![image5](https://img-blog.csdnimg.cn/20181102134026752.png)  
    如果两个都正确，那么登陆成功：  
    ![image6](https://img-blog.csdnimg.cn/20181102134120867.png)  
    在Current.txt中记录当前登陆的用户：  
    ![image7](https://img-blog.csdnimg.cn/201811021342238.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0VtaWx5Qmx1c2U=,size_16,color_FFFFFF,t_70)  

- agenda createMeeting   
    如果创建会议时，参与者不存在：  
    ![image8](https://img-blog.csdnimg.cn/20181102134500641.png)  
    如果创建会议的时间不合法：  
    ![image10](https://img-blog.csdnimg.cn/20181102135723567.png)

    如果参与者在这段时间里没有空： 
    ![image14](https://img-blog.csdnimg.cn/20181102142725677.png) 

    如果成功创建会议，那么也会修改User.json去写入他们创建和参与的会议：  
    ![image13](https://img-blog.csdnimg.cn/20181102142503203.png)  
    我们可以看到在Meeting.json和User.json中都写入了他们参加会议的情况：  
    ![image15](https://img-blog.csdnimg.cn/2018110214290650.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0VtaWx5Qmx1c2U=,size_16,color_FFFFFF,t_70)    

    ![image16](https://img-blog.csdnimg.cn/20181102142930245.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0VtaWx5Qmx1c2U=,size_16,color_FFFFFF,t_70)  


- agenda deleteMeeting  
    如果删除会议的时候会议名字不存在：   
    ![image17](https://img-blog.csdnimg.cn/20181102143039461.png)  

    如果要删除的会议以登入的账号并不是这个会议的发起者的时候：  
    ![image18](https://img-blog.csdnimg.cn/20181102143229879.png)

    成功删除会议，同时我们也会删除会议参与者的相关记录：  
    ![image19](https://img-blog.csdnimg.cn/20181102143328805.png)  
    Meeting.json中记录也将会被删除：  
    ![image20](https://img-blog.csdnimg.cn/20181102143502855.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0VtaWx5Qmx1c2U=,size_16,color_FFFFFF,t_70)  
    User.json中的记录也会被删除：  
    ![image21](https://img-blog.csdnimg.cn/20181102143548770.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0VtaWx5Qmx1c2U=,size_16,color_FFFFFF,t_70)  



- agenda logout  
    成功登出：  
    ![image11](https://img-blog.csdnimg.cn/20181102141715576.png)

- 项目log日志：  
    ![image12](https://img-blog.csdnimg.cn/20181102141817697.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0VtaWx5Qmx1c2U=,size_16,color_FFFFFF,t_70)


