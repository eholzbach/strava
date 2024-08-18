/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CadenceStream struct {
	// The sequence of cadence values for this stream, in rotations per minute
	Data []int32 `json:"data,omitempty"`
	// The number of data points in this stream
	OriginalSize int32 `json:"original_size,omitempty"`
	// The level of detail (sampling) in which this stream was returned
	Resolution string `json:"resolution,omitempty"`
	// The base series used in the case the stream was downsampled
	SeriesType string `json:"series_type,omitempty"`
}
