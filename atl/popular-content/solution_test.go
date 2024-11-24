package popular_content

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstExample(t *testing.T) {
	popularityTracker := NewPopularity()

	popularityTracker.IncreasePopularity(7)
	popularityTracker.IncreasePopularity(7)
	popularityTracker.IncreasePopularity(8)
	assert.Equal(t, 7, int(popularityTracker.GetMostPopular()))
	popularityTracker.IncreasePopularity(8)
	popularityTracker.IncreasePopularity(8)
	assert.Equal(t, 8, int(popularityTracker.GetMostPopular()))
	popularityTracker.DecreasePopularity(8)
	popularityTracker.DecreasePopularity(8)
	assert.Equal(t, 7, int(popularityTracker.GetMostPopular()))
	popularityTracker.DecreasePopularity(7)
	popularityTracker.DecreasePopularity(7)
	popularityTracker.DecreasePopularity(8)
	assert.Equal(t, -1, int(popularityTracker.GetMostPopular()))
}

func TestSecondExample(t *testing.T) {
	p := NewPopularity()
	p.IncreasePopularity(7)
	p.IncreasePopularity(7)
	p.IncreasePopularity(8)
	assert.Equal(t, 7, int(p.GetMostPopular()))
	p.IncreasePopularity(8)
	p.IncreasePopularity(8)
	assert.Equal(t, 8, int(p.GetMostPopular()))
	p.DecreasePopularity(8)
	p.DecreasePopularity(8)
	p.DecreasePopularity(8)
	assert.Equal(t, 7, int(p.GetMostPopular())) // returns 7
	p.DecreasePopularity(7)
	p.DecreasePopularity(7)
	p.DecreasePopularity(7)
	assert.Equal(t, -1, int(p.GetMostPopular())) // returns -1 since there is no content with popularity greater than 0
	assert.Equal(t, 0, len(p.popToContent))
	assert.Equal(t, 0, len(p.contentToPop))
}
