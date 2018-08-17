package interfaceex

import "fmt"

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

func (self User) UpdateNameValue(name string) {
	fmt.Println("[value] Update name to", name)
	self.name = name
}

func (self *User) UpdateNamePointer(name string) {
	fmt.Println("[pointer] Update name to", name)
	self.name = name
}

type Manager struct {
	User
}

//override
func (self *Manager) String() string {
	return fmt.Sprintf("manager %d, %s", self.id, self.name)
}
