import httpClient from './httpClient';

const END_POINT_LOAD_TEST = '/load/test';
const END_POINT_STOP_LOAD = '/stopload';
const END_POINT_LOAD_STATS = '/load/stats';

const startLoadTest = (file, clients, role, cycle, rooms) => httpClient.get(`${END_POINT_LOAD_TEST}?file=${file}&clients=${clients}&role=${role}&cycle=${cycle}&rooms=${rooms}`);

const stopLoadTest = () => httpClient.get(END_POINT_STOP_LOAD);

const loadStats = () => httpClient.get(END_POINT_LOAD_STATS);

export default {
    startLoadTest,
    stopLoadTest,
    loadStats
}