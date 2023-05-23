const minHeap = () => {

    let _data = [];
    
    /**
     * 
     * @param {number} value 
     */
    const insert = (value) => {
        _data.push(value);

        let i = _data.length - 1;
        while (i >= 0) {
            const parentIndex = Math.floor((i - 1) / 2);
            if (_data[parentIndex] > _data[i]) {
                [_data[i], _data[parentIndex]] = [_data[parentIndex], _data[i]];
            }
            i = parentIndex;
        }
    };

    /**
     * 
     * @param {number} value 
     */
    const remove = () => {
        const result = _data[0];
        const last = _data.pop();
        _data[0] = last;

        let i = 0;

        while (i < _data.length) {
            const child1 = i*2+1, child2 = i*2+2;

            if (_data[i] > _data[child1] && (child2 >= _data.length || _data[child1] < _data[child2])) {
                [_data[i], _data[child1]] = [_data[child1], _data[i]];
                i = child1;
            } else if (child2 < _data.length && _data[i] > _data[child2]) {
                [_data[i], _data[child2]] = [_data[child2], _data[i]];
                i = child2;
            } else break;
        }

        return result;
    };
    
    return {
        insert,
        data: _data,
        remove
    };
}

module.exports = { minHeap };