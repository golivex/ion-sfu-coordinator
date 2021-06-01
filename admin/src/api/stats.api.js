import httpClient from './httpClient';

const END_POINT = '/stats';

const getStats = () => httpClient.get(END_POINT);

export default {
    getStats,
}