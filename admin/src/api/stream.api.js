import httpClient from './httpClient';

const END_POINT_START_LIVE_STREAM = '/stream';
const END_POINT_STOP_LIVE_STREAM = '/stopstream';
const END_POINT_DEMO_LIVE_STREAM = '/stream/demo';

const startLiveStream = (session) => httpClient.get(`${END_POINT_START_LIVE_STREAM}/${session}`);

const stopLiveStream = (session) => httpClient.get(`${END_POINT_STOP_LIVE_STREAM}/${session}`)

const demoLiveStream = (session) => httpClient.get(`${END_POINT_DEMO_LIVE_STREAM}/${session}`)


export default {
    startLiveStream,
    stopLiveStream,
    demoLiveStream
}