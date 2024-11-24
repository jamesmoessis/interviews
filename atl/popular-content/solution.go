package popular_content

/*
Imagine you are given a stream of content ids along with an associated action to be performed on them.
Example of contents are video, pages, posts etc. There can be two actions associated with a content id:

increasePopularity → increases the popularity of the content by 1.
The popularity increases when someone comments on the content or likes the content

decreasePopularity → decreases the popularity of the content by 1.
The popularity decreases when a spam bot's/users comments are deleted from the content or its likes are removed from the content

content ids are positive integers

Implement a class that can return the mostPopular content id at any time while consuming the stream of content ids
and its associated action. If there are no contentIds with popularity greater than 0, return -1

interface Popularity {
    void increasePopularity(Integer contentId);
    Integer mostPopular();
    void decreasePopularity(Integer contentId);
}
*/

type ContentID int

type ContentIDSet map[ContentID]bool

type Popularity struct {
	contentToPop map[ContentID]int
	popToContent map[int]ContentIDSet
	maximumPop   int
}

func NewPopularity() Popularity {
	return Popularity{
		contentToPop: make(map[ContentID]int),
		popToContent: make(map[int]ContentIDSet),
		maximumPop:   0,
	}
}

func (m *Popularity) IncreasePopularity(contentID ContentID) {
	// increase pop for content ID
	pop := m.incrementContentPopBy(1, contentID)

	// remove content ID from pop - 1 score
	if toRm, ok := m.popToContent[pop-1]; ok {
		delete(toRm, contentID)
	}

	// Add content ID to score at pop
	if toAdd, ok := m.popToContent[pop]; ok {
		toAdd[contentID] = true
	} else {
		m.popToContent[pop] = make(ContentIDSet)
		m.popToContent[pop][contentID] = true
	}

	if pop > m.maximumPop {
		m.maximumPop = pop
	}
}

func (m *Popularity) DecreasePopularity(contentID ContentID) {
	oldPop, ok := m.contentToPop[contentID]
	if ok {
		m.incrementContentPopBy(-1, contentID)
	}

	if oldContentSet, ok := m.popToContent[oldPop]; ok {
		delete(oldContentSet, contentID)
		// If we decremented the contentID with maximum pop, and it was the only one in the set,
		// we need to also decrement max pop
		if len(oldContentSet) == 0 && oldPop == m.maximumPop {
			m.maximumPop--
		}
		if len(oldContentSet) == 0 {
			delete(m.popToContent, oldPop)
		}
	}

	if oldPop-1 > 0 {
		if toAdd, ok := m.popToContent[oldPop-1]; ok {
			toAdd[contentID] = true
		}
	}
}

func (m *Popularity) GetMostPopular() ContentID {
	contentSet, ok := m.popToContent[m.maximumPop]
	if !ok || m.maximumPop == 0 {
		return -1
	}
	// return first key in set
	for k := range contentSet {
		return k
	}
	return -1
}

func (m *Popularity) incrementContentPopBy(n int, cID ContentID) (newVal int) {
	_, ok := m.contentToPop[cID]
	if !ok {
		m.contentToPop[cID] = n
	} else {
		m.contentToPop[cID] += n
	}
	if m.contentToPop[cID] <= 0 {
		delete(m.contentToPop, cID)
	}
	return m.contentToPop[cID]
}
