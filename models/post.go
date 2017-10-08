package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/russross/blackfriday"
)

// Post 阿诗丹顿啊
type Post struct {
	Title    string    `gorm:"size:60;unique;index;not null" json:"title"`
	Content  string    `gorm:"type:text" json:"content"`
	Public   bool      `gorm:"default:true" json:"-"`
	Tags     []Tag     `gorm:"many2many:blog_post_tags;" json:"tags"`
	Comments []Comment `json:"comments"`
	Cover    string    `json:"cover"`
	gorm.Model
}

// GetPosts a
func GetPosts(page int) interface{} {
	var result []interface{}
	var post []Post
	quantity := 10
	maxLen := 299
	orm.Order("ID desc").Offset((page - 1) * quantity).Limit(quantity).Find(&post)
	for _, v := range post {
		orm.Model(&v).Related(&v.Tags, "Tags").Related(&v.Comments, "Comments")
		contentRune := []rune(v.Content)
		if len(contentRune) > maxLen {
			v.Content = string(contentRune[:maxLen]) + "..."
		}
		v.Content = string(blackfriday.MarkdownCommon([]byte(v.Content)))
		result = append(result, v)
	}
	if result != nil {
		return result
	}
	return []string{}
}

// GetPost ..
func GetPost(id int) (Post, error) {
	var post Post
	ret := orm.First(&post, id).Model(&post).Related(&post.Tags, "Tags").Related(&post.Comments, "Comments")
	if ret.Error != nil {
		return Post{}, ret.Error
	}
	return post, nil
}

// GetPostByTitle ..
func GetPostByTitle(title string) (Post, error) {
	var result []Comment
	var post Post
	var ret = orm.Where("Title = ?", title).First(&post)
	ret = orm.Model(&post).Related(&post.Tags, "Tags").Related(&post.Comments, "Comments")
	if ret.Error != nil {
		return post, ret.Error
	}

	for _, c := range post.Comments {
		orm.Model(&c).Related(&c.User, "User")
		result = append(result, c)
	}
	if len(result) > 0 {
		post.Comments = result
	}

	markdown := []byte(post.Content)
	html := blackfriday.MarkdownCommon(markdown)

	// highlighted, err := syntaxhighlight.AsHTML(html)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	post.Content = string(html)
	return post, nil
}

// NewPost ..
func NewPost(post Post) (Post, error) {
	if len(post.Title) > 0 && len(post.Content) > 0 {
		ret := orm.Create(&post)
		if ret.Error != nil {
			return post, ret.Error
		}
	} else {
		return post, errors.New("参数错误")
	}
	return post, nil
}

// PutPost ..
func PutPost(ID int, p Post) (Post, error) {
	var post Post
	ret := orm.First(&post, ID)
	if ret.Error != nil {
		return post, ret.Error
	}

	if p.Title != "" {
		post.Title = p.Title
	}
	if p.Content != "" {
		post.Content = p.Content
	}
	// post.Tags = p.Content

	ret = orm.Save(&post)
	if ret.Error != nil {
		return post, ret.Error
	}
	return post, nil
}

// DeletePost ..
func DeletePost(ID int) (Post, error) {
	p, err := GetPost(ID)
	if err != nil {
		return p, err
	}
	orm.Delete(&p)
	return p, nil
}
