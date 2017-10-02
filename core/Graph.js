// this module is used to create the graph

const _ = require('lodash');
const Node = require('./Node');

class Graph {
	constructor() {
		this.nodes = {}; // contains list of nodes, for quick lookup
		this.length = 0;
	}
	_hasNode(id) {
		return !!this.nodes[id];
	}
	addNode(id, params) {
		if(this._hasNode(id)) {
			throw new Error('Node ID already exists, Specify new id');
		} else {
			this.nodes[id] = new Node(id, params);
		}
	}
	addEdge(from, to, params) {
		if(!this._hasNode(from) || !this._hasNode(to)) {
			throw new Error('Invalid Source or destination nodes');
		} else {
			this.nodes[from].addNeighbour(to, params);
		}
	}
	removeEdge(from, to) {
		if(!this._hasNode(from) || !this._hasNode(to)) {
			throw new Error('Invalid Source or destination nodes');
		} else {
			this.nodes[from].removeNeighbour(to);
		}
	}
	hasAnEdge(from, to) {
		return this.nodes[from].isNeighbourOf(to);
	}
}
