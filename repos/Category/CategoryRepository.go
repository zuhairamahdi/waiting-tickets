package category

import "github.com/zuhairamahdi/gin-test/entities"

//InsertCategory to Insert new Category
func InsertCategory(category entities.Category) {
	//TODO Insert Category Here
}

//DeleteCategory to delete a category
func DeleteCategory(id int) {
	//TODO delete category from database here
}

//UpdateCategory to update a category
func UpdateCategory(id int, newName string) {
	//TODO Update category here
}

//UpdateCategoryWithObject
func UpdateCategoryWithObject(id int, category entities.Category) {
	//TODO Update category using category type object
}
