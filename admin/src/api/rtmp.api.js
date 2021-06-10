import httpClient from './httpClient';

const END_POINT_START_LIVE_RTMP = '/rtmp';
const END_POINT_STOP_LIVE_RTMP = '/stoprtmp';
const END_POINT_DEMO_LIVE_RTMP = '/rtmp/demo';

const startLiveRtmp = (session) => httpClient.get(`${END_POINT_START_LIVE_RTMP}/${session}`);

const stopLiveRtmp = (session) => httpClient.get(`${END_POINT_STOP_LIVE_RTMP}/${session}`);

const demoLiveRtmp = (session) => httpClient.get(`${END_POINT_DEMO_LIVE_RTMP}/${session}`)


export default {
    startLiveRtmp,
    stopLiveRtmp,
    demoLiveRtmp
}