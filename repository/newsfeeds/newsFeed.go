package newsfeeds

type Item struct {
	Title string `json: "title"`
	Body string   `json: "body"`
}


type NewsFeeds struct{
	Items []Item 
}


//func function_name  return_type
func New() (n *NewsFeeds){
	return &NewsFeeds{}
}

//func Input_type function_name
func (r *NewsFeeds) AddItem(item Item) {
	r.Items=append(r.Items, item)
}
//func Input_type function_name return_type
func (r *NewsFeeds) GETAll() []Item {
	return r.Items;
}