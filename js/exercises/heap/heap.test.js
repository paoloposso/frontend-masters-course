const { minHeap } = require('./heap');

test('Test add to heap', () => {
    let heap = minHeap();
    heap.insert(1);

    expect(heap.data[0]).toBe(1);
});

test('Test add multiple to heap', () => {
    let heap = minHeap();

    heap.insert(30);
    heap.insert(15);
    heap.insert(100);
    heap.insert(50);
    heap.insert(75);
    heap.insert(10);
    heap.insert(90);
  
    expect(heap.data[0]).toBe(10);
});

test('Test remove one from heap', () => {
    let heap = minHeap();

    heap.insert(30);
    heap.insert(15);
    heap.insert(100);

    expect(heap.remove()).toBe(15);
    expect(heap.data[0]).toBe(30);
});