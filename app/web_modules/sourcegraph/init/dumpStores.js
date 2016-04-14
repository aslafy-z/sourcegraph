import BlobStore from "sourcegraph/blob/BlobStore";
import DefStore from "sourcegraph/def/DefStore";
import RepoStore from "sourcegraph/repo/RepoStore";
import TreeStore from "sourcegraph/tree/TreeStore";
import DashboardStore from "sourcegraph/dashboard/DashboardStore";
import BuildStore from "sourcegraph/build/BuildStore";
import EventLogger from "sourcegraph/util/EventLogger";

export default function dumpStores() {
	return {
		RepoStore,
		BlobStore,
		DefStore,
		TreeStore,
		DashboardStore,
		BuildStore,
		EventLogger,
	};
}