// This file was auto-generated by Fern from our API Definition.

package cloud

type GetRayPerfLogsResponse struct {
	// A list of service performance summaries.
	PerfLists []*SvcPerf `json:"perf_lists,omitempty"`
}
