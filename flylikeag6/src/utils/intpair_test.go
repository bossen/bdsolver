package utils

import ( 
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestIsIntPairInSlice(t *testing.T) {
	intpairslice := []IntPair{IntPair{0, 1}, IntPair{4, 1}, IntPair{7, 3}, IntPair{2, 5}}
	
	assert.True(t, IsIntPairInSlice(IntPair{0, 1}, intpairslice), "the IntPair was not found")
	assert.True(t, IsIntPairInSlice(IntPair{4, 1}, intpairslice), "the IntPair was not found")
	assert.True(t, IsIntPairInSlice(IntPair{7, 3}, intpairslice), "the IntPair was not found")
	assert.True(t, IsIntPairInSlice(IntPair{2, 5}, intpairslice), "the IntPair was not found")
	
	assert.False(t, IsIntPairInSlice(IntPair{1, 5}, intpairslice), "the IntPair was found")
	assert.False(t, IsIntPairInSlice(IntPair{1, 0}, intpairslice), "the IntPair was found")
	assert.False(t, IsIntPairInSlice(IntPair{9, 9}, intpairslice), "the IntPair was found")
}
