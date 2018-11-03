``` go
package service

//用户注册
func UserRegister(userName string, password string,
  email string, phone string) error {
  return nil
}

//用户登录
func UserLogin(userName string, password string) error {
  return nil
}

//用户登出
func UserLogout() error {
  return nil
}

//查询所有用户
func QueryAllUsers() []model.User {
  return []model.User{}
}


//删除用户
//删除该用户的用户信息以及参加和发起的会议，如果该用户是某一个会议的发起者，则删除该会议，
//如果由于删除该用户造成某一个会议的参与者变成0，则删除该会议
func DeleteUser(password string) error {
  return nil
}


//创建会议
func CreateMeeting(title string, startDate string,
  endDate string, participators []string) error {
  return nil
}

//添加自己发起的某一会议的一个参与者
func AddParticipator(title string, paticipator string) error {
  return nil
}

//删除自己创建的某一会议的一个参与者
//如果会议的参与者因此变成0，则删除该会议
func DeleteParticipator(title string, participator string) error {
  return nil
}

//查询会议，通过开始时间和结束时间查询当前用户需要参加的所有会议
func QueryMeeting(startDate string, endDate string) ([]model.Metting, error) {
  return []model.Meeting{},nil
}

//取消会议
func DeleteMeeting(title string) error {
  return nil
}

//退出会议
//如果退出会议之后的参与者为0的会议将会被删除
func QuitMeeting(title string) error {
  return nil
}

//清空当前用户发起的所有会议安排
func DeleteAllMeeting() error {
  return nil
}


```
