package models

type Item struct {
	Id    string
	Price int
}

func (i *Item) Constructor(Id string, Price int) {
	i.Id = Id
	i.Price = Price
}

/**
* get the id of an item
* @return id item
 */
func (i *Item) GetId() string {
	return i.Id
}

/**
* assigns the id of an item
* @param id item
 */
func (i *Item) SetId(Id string) {
	i.Id = Id
}

/**
* get the Price of an item
* @return Price item
 */
func (i *Item) GetPrice() int {
	return i.Price
}

/**
* assigns the Price of an item
* @param Price item
 */
func (i *Item) SetPrice(Price int) {
	i.Price = Price
}

type ResponseItems struct {
	Code int
	Body Item
}
