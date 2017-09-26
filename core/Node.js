class Node {
	constructor(id, params = {}) {
		this.id = id;
		this.params = params;
		this.neighbours = {};
		this.neighbourCount = 0;
	}
	getNode() {
		return {
			id: this.id,
			...this.params,
		};
	}
	addNeighbour(id, params = {}) {
		// since, it's a generic method hence, we are just adding a id to list of this node
		if(!this.neighbours[id]) {
			this.neighbours[id] = {
				id,
				...params,
			};
		}
		this.neighbourCount++;
		return this;
	}
	removeNeighbour(id) {
		this.neighbours[id] = undefined;
		this.neighbourCount--;
	}
	getNeighbourSet() {
		const neighborArray = [];
		for(const neighbor of this.neighbours) {
			neighborArray.push(neighbor);
		}
		return neighborArray;
	}
}

module.exports = Node;