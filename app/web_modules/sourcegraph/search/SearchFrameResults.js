import React from "react";

import Component from "sourcegraph/Component";

class SearchFrameResultsView extends Component {

	reconcileState(state, props) {
		Object.assign(state, props);
		if (props.resultData.Error) {
			state.error = props.resultData.Error;
			return;

		}
		state.html = props.resultData.HTML;
		state.total = props.resultData.Total;
	}

	render() {
		if (this.state.error) {
			return (<div className="alert alert-warning">There was an error returning your results: {this.state.error}</div>);
		}

		let summary;
		if (this.state.total === 0) {
			summary = `No results found for "${this.state.query}"`;
		} else {
			let s = this.state.total === 1 ? "" : "s";
			summary = `${this.state.total} result${s} for "${this.state.query}"`;
			if (this.state.currentPage > 1) summary = `Page ${this.state.currentPage} of ${summary}`;
		}

		return (
			<div className={`${this.state.label}-search-results`}>
				<p className="summary">{summary}</p>
				<div dangerouslySetInnerHTML={{__html: this.state.html}} />
			</div>
			);
	}
}


SearchFrameResultsView.propTypes = {
	repo: React.PropTypes.string,
	query: React.PropTypes.string,
	rev: React.PropTypes.string,
	page: React.PropTypes.number,
	label: React.PropTypes.string,
	resultData: React.PropTypes.object,
};

export default SearchFrameResultsView;