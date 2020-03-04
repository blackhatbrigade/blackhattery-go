package blogs

import (
	"github.com/jetbasrawi/go.cqrs"
	"github.com/jetbasrawi/go.geteventstore"
)

type Repo struct {
	repo *ycq.GetEventStoreCommonDomainRepo
}

func NewBlogRepo(eventStore *goes.Client, eventBus ycq.EventBus) (*Repo, error) {

	r, err := ycq.NewCommonDomainRepository(eventStore, eventBus)
	if err != nil {
		return nil, err
	}

	ret := &Repo{
		repo: r,
	}

	// An aggregate factory creates an aggregate instance given the name of an aggregate.
	aggregateFactory := ycq.NewDelegateAggregateFactory()
  aggregateFactory.RegisterDelegate(&BlogPost{},
		func(id string) ycq.AggregateRoot { return NewBlogPost(id) })
/*	ret.repo.SetAggregateFactory(aggregateFactory)

	// A stream name delegate constructs a stream name.
	// A common way to construct a stream name is to use a bounded context and
	// an aggregate id.
	// The interface for a stream name delegate takes a two strings. One may be
	// the aggregate type and the other the aggregate id. In this case the first
	// argument and the second argument are concatenated with a hyphen.
	streamNameDelegate := ycq.NewDelegateStreamNamer()
	streamNameDelegate.RegisterDelegate(func(t string, id string) string {
		return t + "-" + id
	}, &InventoryItem{})
	ret.repo.SetStreamNameDelegate(streamNameDelegate)

	// An event factory creates an instance of an event given the name of an event
	// as a string.
	eventFactory := ycq.NewDelegateEventFactory()
	eventFactory.RegisterDelegate(&InventoryItemCreated{},
		func() interface{} { return &InventoryItemCreated{} })
	eventFactory.RegisterDelegate(&InventoryItemRenamed{},
		func() interface{} { return &InventoryItemRenamed{} })
	eventFactory.RegisterDelegate(&InventoryItemDeactivated{},
		func() interface{} { return &InventoryItemDeactivated{} })
	eventFactory.RegisterDelegate(&ItemsRemovedFromInventory{},
		func() interface{} { return &ItemsRemovedFromInventory{} })
	eventFactory.RegisterDelegate(&ItemsCheckedIntoInventory{},
		func() interface{} { return &ItemsCheckedIntoInventory{} })
	ret.repo.SetEventFactory(eventFactory)
  */

	return ret, nil
}
