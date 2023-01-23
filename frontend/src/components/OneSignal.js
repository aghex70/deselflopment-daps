import {useEffect} from 'react';
import OneSignal from 'react-onesignal';

const OneSignalNotifier = () => {
    useEffect(() => {
        OneSignal.init({
            appId: "d70c53bb-aad1-461e-96a6-b6cec7b9a0d4",
        });
    }, []);
}

export default OneSignalNotifier;