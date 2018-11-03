# Agenda

## contributors:                          
16340270 杨捷
16340272 杨淼   

## Usage                                
　　`agd command`

- 注册账户      
　　`agd regist -u usernaeme -p password -e e-mail -t telephone`     

- 登陆  
　　`agd login -u username -p password`

- 登出   
　　`agd logout`

- 显示所有用户    
　　`agd lsu`

- 删除当前用户        
　　`agd del -p password`

- 创建会议                      
　　`agd cm -t title -p participator1 -p participator2 -s start -e end`

- 添删会议成员                     
　　add: `agd ap -t title -p name`                                   
　　delete: `agd dp -t title -p name`

- 显示一个时间段内的所有会议                  
　　`agd lsm -s start -e end`

- 取消某个会议                 
　　`agd cancel -t title`

- 退出某个会议                
　　`agd quit -t title`

- 清楚所有会议                 
　　`agd clear`

## log & data
- logs                  
　　存储所有账户信息

- dates                           
　　存储所有会议相关信息


## examples


```
./agd regist -u user1 -p 123 -e email1@mail.com -t 11111
Success

./agd regist -u user2 -p 122 -e email2@mail.com -t 22222
Success

./agd regist -u user3 -p 333 -e email3@mail.com -t 33333
Success

./agd regist -u user1 -p 111 -e mail2@mail.com -t 11111
this username is aleardy exist

./agd login -u usr1 -p 123
username not exist

./agd login -u user1 -p 12
incorrect password

./agd login -u user1 -p 123
Success

./agd logout
Success

./agd lsu
user1, email1@mail.com, 1111
user2, email2@mail.com, 2222
user3, email3@mail.com, 3333

./agd lsu
user1, email1@mail.com, 1111
user2, email2@mail.com, 2222
user3, email3@mail.com, 3333

./agd login -u user1 -p 123
Success

./agd del -p 124
incorrect password

./agd del -p 123
Success

./agd lsu
user2, email2@mail.com, 2222
user3, email3@mail.com, 3333

./agd cm -t title -p user1, user2 -s 2018-11-01T10:00 -e 2018-11-01T10:30
no user login

./agd login -u user3 -p333
Success

./agd cm -t title -p user1, user2 -s 2018-11-01T10:00 -e 2018-11-01T10:30
Success

./agd lsm -s 2018-11-01T10:00 -e 2018-11-01T10:30
title
　- sponsor: user3
　- time: 2018-11-01 10:00 - 2018-11-01 10:30
　- participators: user1

./agd lsm -s 2018-11-01T0:00 -e 2018-11-01T20:00
title
　- sponsor: user3
　- time: 2018-11-01 10:00 - 2018-11-01 10:30
　- participators: user1

title2
　- sponsor: user3
　- time: 2018-11-01 11:00 - 2018-11-01 11:30
　- participators: user1

title3
　- sponsor: user3
　- time: 2018-11-01 11:00 - 2018-11-01 11:30
　- participators: user2

./agd ap -t title2 -p user
participator user is not exist

./agd ap -t title2 -p user2
Success

./agd dp -t title1 -p user2
the title not exist

./agd dp -t title3 -p user2
Success

./agd lsm -s 2018-11-01T0:00 -e 2018-11-01T23:00
title
　- sponsor: user3
　- time: 2018-11-01 10:00 - 2018-11-01 10:30
　- participators: user1

title2
　- sponsor: user3
　- time: 2018-11-01 11:00 - 2018-11-01 11:30
　- participators: user1, user2

./agd cancel -t title
Success

./agd lsm -s 2018-11-01T0:00 -e 2018-11-01T23:00
title2
　- sponsor: user3
　- time: 2018-11-01 11:00 - 2018-11-01 11:30
　- participators: user1, user2

./agd lsm -s 2018-11-01T0:00 -e 2018-11-01T23:00
title3
　- sponsor: user3
　- time: 2018-11-01 11:00 - 2018-11-01 11:30
　- participators: user2

title1
　- sponsor: user2
　- time: 2018-11-01 13:00 - 2018-11-01 15:30
　- participators: user1, user3

title2
　- sponsor: user2
　- time: 2018-11-01 13:00 - 2018-11-01 15:30
　- participators: user1


./agd quit -t title1
current user is the sponsor of the meeting, can't quit

./agd quit -t title3
Success

./agd lsm -s 2018-11-01T0:00 -e 2018-11-01T23:00
title1
　- sponsor: user2
　- time: 2018-11-01 13:00 - 2018-11-01 15:30
　- participators: user1, user3

title2
　- sponsor: user2
　- time: 2018-11-01 13:00 - 2018-11-01 15:30
　- participators: user1

./agd lsm -s 2018-11-01T0:00 -e 2018-11-01T23:00
title2
　- sponsor: user3
　- time: 2018-11-01 11:00 - 2018-11-01 11:30
　- participators: user1, user2

title
　- sponsor: user3
　- time: 2018-11-01 11:00 - 2018-11-01 11:30
　- participators: user1

title3
　- sponsor: user3
　- time: 2018-11-01 11:00 - 2018-11-01 11:30
　- participators: user2

./agd clear
Success

./agd lsm -s 2018-11-01T0:00 -e 2018-11-01T23:00


```
