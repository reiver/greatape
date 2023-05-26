package core

import (
	"fmt"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/contracts/model"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/security"
)

type categoryType struct {
	object
	description string
}

// noinspection GoUnusedExportedFunction
func InitializeCategoryType() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewCategoryType(id int64, description string) (ICategoryType, error) {
	instance := &categoryType{
		object: object{
			id: id,
		},
		description: description,
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewCategoryTypeFromEntity(entity ICategoryTypeEntity) (ICategoryType, error) {
	instance := &categoryType{
		object: object{
			id: entity.Id(),
		},
		description: entity.Description(),
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (categoryType *categoryType) Description() string {
	return categoryType.description
}

func (categoryType *categoryType) UpdateDescription(description string, editor Identity) {
	if err := repository.CategoryTypes.UpdateDescription(categoryType.id, description, editor.Id()); err != nil {
		panic(err.Error())
	}

	categoryType.description = description
}

func (categoryType *categoryType) UpdateDescriptionAtomic(transaction ITransaction, description string, editor Identity) {
	transaction.OnCommit(func() {
		categoryType.description = description
	})

	if err := repository.CategoryTypes.UpdateDescriptionAtomic(transaction, categoryType.id, description, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (categoryType *categoryType) Validate() error {
	return nil
}

func (categoryType *categoryType) String() string {
	return fmt.Sprintf("CategoryType (Id: %d, Description: %v)", categoryType.Id(), categoryType.Description())
}

//------------------------------------------------------------------------------

type categoryTypes struct {
	collection CategoryTypes
}

// NewCategoryTypes creates an empty collection of 'Category Type' which is not thread-safe.
func NewCategoryTypes() ICategoryTypeCollection {
	return &categoryTypes{
		collection: make(CategoryTypes, 0),
	}
}

func (categoryTypes *categoryTypes) Count() int {
	return len(categoryTypes.collection)
}

func (categoryTypes *categoryTypes) IsEmpty() bool {
	return len(categoryTypes.collection) == 0
}

func (categoryTypes *categoryTypes) IsNotEmpty() bool {
	return len(categoryTypes.collection) > 0
}

func (categoryTypes *categoryTypes) HasExactlyOneItem() bool {
	return len(categoryTypes.collection) == 1
}

func (categoryTypes *categoryTypes) HasAtLeastOneItem() bool {
	return len(categoryTypes.collection) >= 1
}

func (categoryTypes *categoryTypes) First() ICategoryType {
	return categoryTypes.collection[0]
}

func (categoryTypes *categoryTypes) Append(categoryType ICategoryType) {
	categoryTypes.collection = append(categoryTypes.collection, categoryType)
}

func (categoryTypes *categoryTypes) Reverse() ICategoryTypeCollection {
	slice := categoryTypes.collection

	start := 0
	end := len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}

	categoryTypes.collection = slice

	return categoryTypes
}

func (categoryTypes *categoryTypes) ForEach(iterator CategoryTypeIterator) {
	if iterator == nil {
		return
	}

	for _, value := range categoryTypes.collection {
		iterator(value)
	}
}

func (categoryTypes *categoryTypes) Array() CategoryTypes {
	return categoryTypes.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) CategoryTypeExists(id int64) bool {
	return dispatcher.conductor.CategoryTypeManager().Exists(id)
}

func (dispatcher *dispatcher) CategoryTypeExistsWhich(condition CategoryTypeCondition) bool {
	return dispatcher.conductor.CategoryTypeManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListCategoryTypes() ICategoryTypeCollection {
	return dispatcher.conductor.CategoryTypeManager().ListCategoryTypes(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachCategoryType(iterator CategoryTypeIterator) {
	dispatcher.conductor.CategoryTypeManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterCategoryTypes(predicate CategoryTypeFilterPredicate) ICategoryTypeCollection {
	return dispatcher.conductor.CategoryTypeManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapCategoryTypes(predicate CategoryTypeMapPredicate) ICategoryTypeCollection {
	return dispatcher.conductor.CategoryTypeManager().Map(predicate)
}

func (dispatcher *dispatcher) GetCategoryType(id int64) ICategoryType {
	if categoryType, err := dispatcher.conductor.CategoryTypeManager().GetCategoryType(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return categoryType
	}
}

func (dispatcher *dispatcher) AddCategoryType(description string) ICategoryType {
	transaction := dispatcher.transaction
	if transaction != nil {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().AddCategoryTypeAtomic(transaction, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	} else {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().AddCategoryType(description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	}
}

func (dispatcher *dispatcher) AddCategoryTypeWithCustomId(id int64, description string) ICategoryType {
	transaction := dispatcher.transaction
	if transaction != nil {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().AddCategoryTypeWithCustomIdAtomic(id, transaction, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	} else {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().AddCategoryTypeWithCustomId(id, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	}
}

func (dispatcher *dispatcher) LogCategoryType(description string, source string, payload string) {
	dispatcher.conductor.CategoryTypeManager().Log(description, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateCategoryType(id int64, description string) ICategoryType {
	transaction := dispatcher.transaction
	if transaction != nil {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().UpdateCategoryTypeAtomic(transaction, id, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	} else {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().UpdateCategoryType(id, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateCategoryTypeObject(object IObject, categoryType ICategoryType) ICategoryType {
	transaction := dispatcher.transaction
	if transaction != nil {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().UpdateCategoryTypeAtomic(transaction, object.Id(), categoryType.Description(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	} else {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().UpdateCategoryType(object.Id(), categoryType.Description(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateCategoryTypeObject(object IObject, categoryType ICategoryType) ICategoryType {
	transaction := dispatcher.transaction
	if transaction != nil {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().AddOrUpdateCategoryTypeObjectAtomic(transaction, object.Id(), categoryType, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	} else {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().AddOrUpdateCategoryTypeObject(object.Id(), categoryType, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	}
}

func (dispatcher *dispatcher) RemoveCategoryType(id int64) ICategoryType {
	transaction := dispatcher.transaction
	if transaction != nil {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().RemoveCategoryTypeAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	} else {
		if categoryType, err := dispatcher.conductor.CategoryTypeManager().RemoveCategoryType(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return categoryType
		}
	}
}
