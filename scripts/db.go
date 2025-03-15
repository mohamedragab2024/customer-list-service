package main

import (
	"customer-list-service/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func insertDummyData() {
	customers := []models.Customer{
		{Name: "John Doe", Address: "123 Main St", Status: "Active", Email: "john@example.com"},
		{Name: "Jane Smith", Address: "456 Elm St", Status: "Inactive", Email: "jane@example.com"},
		{Name: "Bob Johnson", Address: "789 Oak St", Status: "Active", Email: "bob@example.com"},
		{Name: "Alice Brown", Address: "101 Pine St", Status: "Active", Email: "alice@example.com"},
		{Name: "Charlie Davis", Address: "202 Maple St", Status: "Inactive", Email: "charlie@example.com"},
		{Name: "Eve White", Address: "303 Birch St", Status: "Active", Email: "eve@example.com"},
		{Name: "Frank Black", Address: "404 Cedar St", Status: "Inactive", Email: "frank@example.com"},
		{Name: "Grace Green", Address: "505 Walnut St", Status: "Active", Email: "grace@example.com"},
		{Name: "Hank Blue", Address: "606 Chestnut St", Status: "Inactive", Email: "hank@example.com"},
		{Name: "Ivy Red", Address: "707 Spruce St", Status: "Active", Email: "ivy@example.com"},
		{Name: "Jack Yellow", Address: "808 Fir St", Status: "Inactive", Email: "jack@example.com"},
		{Name: "Kathy Purple", Address: "909 Ash St", Status: "Active", Email: "kathy@example.com"},
		{Name: "Leo Orange", Address: "1010 Poplar St", Status: "Inactive", Email: "leo@example.com"},
		{Name: "Mona Pink", Address: "1111 Willow St", Status: "Active", Email: "mona@example.com"},
		{Name: "Nina Gray", Address: "1212 Redwood St", Status: "Inactive", Email: "nina@example.com"},
		{Name: "Oscar Silver", Address: "1313 Cypress St", Status: "Active", Email: "oscar@example.com"},
		{Name: "Paul Gold", Address: "1414 Palm St", Status: "Inactive", Email: "paul@example.com"},
		{Name: "Quincy Bronze", Address: "1515 Magnolia St", Status: "Active", Email: "quincy@example.com"},
		{Name: "Rita Copper", Address: "1616 Dogwood St", Status: "Inactive", Email: "rita@example.com"},
		{Name: "Sam Steel", Address: "1717 Hickory St", Status: "Active", Email: "sam@example.com"},
		{Name: "Tina Brass", Address: "1818 Sycamore St", Status: "Inactive", Email: "tina@example.com"},
		{Name: "Uma Zinc", Address: "1919 Beech St", Status: "Active", Email: "uma@example.com"},
		{Name: "Vince Iron", Address: "2020 Hemlock St", Status: "Inactive", Email: "vince@example.com"},
		{Name: "Wendy Lead", Address: "2121 Alder St", Status: "Active", Email: "wendy@example.com"},
		{Name: "Xander Tin", Address: "2222 Sequoia St", Status: "Inactive", Email: "xander@example.com"},
		{Name: "Yara Nickel", Address: "2323 Juniper St", Status: "Active", Email: "yara@example.com"},
		{Name: "Zane Cobalt", Address: "2424 Cedar St", Status: "Inactive", Email: "zane@example.com"},
		{Name: "Amy Quartz", Address: "2525 Birch St", Status: "Active", Email: "amy@example.com"},
		{Name: "Brian Topaz", Address: "2626 Maple St", Status: "Inactive", Email: "brian@example.com"},
		{Name: "Cathy Ruby", Address: "2727 Pine St", Status: "Active", Email: "cathy@example.com"},
		{Name: "David Sapphire", Address: "2828 Elm St", Status: "Inactive", Email: "david@example.com"},
		{Name: "Ella Emerald", Address: "2929 Oak St", Status: "Active", Email: "ella@example.com"},
		{Name: "Fred Diamond", Address: "3030 Fir St", Status: "Inactive", Email: "fred@example.com"},
		{Name: "Gina Pearl", Address: "3131 Spruce St", Status: "Active", Email: "gina@example.com"},
		{Name: "Hugo Jade", Address: "3232 Cedar St", Status: "Inactive", Email: "hugo@example.com"},
		{Name: "Iris Opal", Address: "3333 Chestnut St", Status: "Active", Email: "iris@example.com"},
		{Name: "Jake Amber", Address: "3434 Walnut St", Status: "Inactive", Email: "jake@example.com"},
		{Name: "Kara Garnet", Address: "3535 Ash St", Status: "Active", Email: "kara@example.com"},
		{Name: "Liam Onyx", Address: "3636 Poplar St", Status: "Inactive", Email: "liam@example.com"},
		{Name: "Mia Coral", Address: "3737 Willow St", Status: "Active", Email: "mia@example.com"},
		{Name: "Noah Jasper", Address: "3838 Redwood St", Status: "Inactive", Email: "noah@example.com"},
		{Name: "Olivia Agate", Address: "3939 Cypress St", Status: "Active", Email: "olivia@example.com"},
		{Name: "Pete Lapis", Address: "4040 Palm St", Status: "Inactive", Email: "pete@example.com"},
		{Name: "Quinn Beryl", Address: "4141 Magnolia St", Status: "Active", Email: "quinn@example.com"},
		{Name: "Rose Malachite", Address: "4242 Dogwood St", Status: "Inactive", Email: "rose@example.com"},
		{Name: "Steve Turquoise", Address: "4343 Hickory St", Status: "Active", Email: "steve@example.com"},
		{Name: "Tara Amethyst", Address: "4444 Sycamore St", Status: "Inactive", Email: "tara@example.com"},
		{Name: "Umar Citrine", Address: "4545 Beech St", Status: "Active", Email: "umar@example.com"},
		{Name: "Vera Peridot", Address: "4646 Hemlock St", Status: "Inactive", Email: "vera@example.com"},
		{Name: "Will Garnet", Address: "4747 Alder St", Status: "Active", Email: "will@example.com"},
		{Name: "Xena Zircon", Address: "4848 Sequoia St", Status: "Inactive", Email: "xena@example.com"},
		{Name: "Yusuf Moonstone", Address: "4949 Juniper St", Status: "Active", Email: "yusuf@example.com"},
		{Name: "Zara Sunstone", Address: "5050 Cedar St", Status: "Inactive", Email: "zara@example.com"},
	}

	for _, customer := range customers {
		models.DB.Create(&customer)
	}
}

func main() {
	var err error
	models.DB, err = gorm.Open(sqlite.Open("customers.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	models.DB.AutoMigrate(&models.Customer{})
	insertDummyData()
}
