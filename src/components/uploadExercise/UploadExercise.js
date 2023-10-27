import { UploadOutlined } from '@ant-design/icons';
import { Button, message, Space, Upload } from 'antd';
import axios from 'axios';
import { useContext, useState } from 'react';
import { CourseContext } from '../course/Course';

export const UploadExercise = ({ exerciseId }) => {
    const [uploaded, setUploaded] = useState(false);
    const courseId = useContext(CourseContext);

    const customRequest = async ({ onSuccess, onError, file }) => {
        try {

            const values = {
                "exerciseId": exerciseId,
                "courseId": courseId,
                "file": file.uid
            }

            const response = await axios.post(`http://localhost:4000/submitted${courseId}`, values);

            if (response.status === 201) {
                message.success(`${file.name} file uploaded successfully`);
                onSuccess();
                setUploaded(true);
            }
            else {
                message.error(`${file.name} file upload failed.`);
                onError(response.statusText);
            }

        }
        catch (error) {
            message.error(`${file.name} file upload failed.`);
            onError(error);
        }
    }

    const onRemove = (file) => {
        setUploaded(false);
        axios
            .delete(`http://localhost:4000/submitted${courseId}/${exerciseId}`)
            .then((response) => {
                message.success(`${file.name} file removed successfully`);
            })
            .catch((error) => {
                message.error(`${file.name} file remove failed.`);
            });
    }
    return (
        <Space
            direction="vertical"
            style={{
                width: '100%',
            }}
            size="large"
        >
            <Upload
                customRequest={customRequest}
                onRemove={onRemove}
                listType="picture"
                maxCount={1}
            >
                <Button disabled={uploaded} icon={<UploadOutlined />}>Click to Upload Exercise</Button>
            </Upload>
        </Space >
    )
}
