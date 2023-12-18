import React from 'react'
import { Button, message } from 'antd';

export const DownloadExercise = ({ name, url }) => {
    const downloadExerciseFile = () => {

        if (!url) {
            message.error('Exercise file not found.');
            return;
        }

        const a = document.createElement('a');
        a.href = url;
        a.download = name;
        document.body.appendChild(a);
        a.click();
        window.URL.revokeObjectURL(url);
    };

    return (
        <Button onClick={downloadExerciseFile}>Download Exercise File</Button>
    )
}
