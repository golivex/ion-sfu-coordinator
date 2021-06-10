import httpClient from './httpClient';

const END_POINT_START_DISK_SAVE = '/disk';
const END_POINT_STOP_DISK_SAVE = '/stopdisk';

const startDiskSave = (session, filename) => httpClient.get(`${END_POINT_START_DISK_SAVE}/${session}/${filename}`);

const stopDiskSave = (session) => httpClient.get(`${END_POINT_STOP_DISK_SAVE}/${session}`);


export default {
    startDiskSave,
    stopDiskSave
}