import { useState, useEffect, useContext } from 'react';
import { Checkbox, Upload, Button, message } from 'antd';
import { UploadOutlined } from '@ant-design/icons';
import { useUser } from '../../contexts/UserContext';
import { useAuth } from '../../contexts/AuthContext';
import { CourseContext } from '../course/Course';

export const Exercises = ({ id, question, topicId }) => {
    const [completed, setCompleted] = useState(false);
    const [isSetLoading, setIsSetLoading] = useState(true);
    const [defaultFiles, setDefaultFiles] = useState([])
    const { currentUser } = useAuth();
    const { setsubmittedInLocalStorage, getSubmittedFromLocalStorage, removeFromsubmittedInLocalStorage } = useUser();
    const courseId = useContext(CourseContext);

    const handleCheckboxChange = (e) => {
        setCompleted(e.target.checked);
        removeFromsubmittedInLocalStorage(currentUser.email, courseId, topicId, id);
    };

    const customRequest = async ({ onSuccess, onError, file }) => {
        try {
            const reader = new FileReader();
            reader.onloadend = () => {
                const data = reader.result;
                setsubmittedInLocalStorage(currentUser.email, courseId, topicId, id, file.name, data)
                setCompleted(true);
                onSuccess();
                message.success(`${file.name} uploaded successfully.`);
            };
            reader.readAsDataURL(file);
        }
        catch (error) {
            message.error(`${file.name} file upload failed.`);
            onError(error);
        }
    };

    const onRemove = (file) => {
        setCompleted(false);
        message.success(`${file.name} removed successfully.`);
        setDefaultFiles([{}]);
        console.log(defaultFiles)
        return removeFromsubmittedInLocalStorage(currentUser.email, courseId, topicId, id);
    }

    useEffect(() => {
        setIsSetLoading(true)
        const storedData = getSubmittedFromLocalStorage(currentUser.email, courseId, topicId, id);
        if (storedData) {
            const key = `exercise-${courseId}-${topicId}-${id}`;
            setDefaultFiles([{
                name: storedData.fileName,
                thumbUrl: storedData[key],
                status: 'done',
            }])
            setCompleted(true);
        }
        setIsSetLoading(false)
    }, [completed, getSubmittedFromLocalStorage, currentUser.email, courseId, topicId, id, setDefaultFiles]);

    return (
        <>
            <Checkbox className="size mx-5" disabled={!completed} checked={completed} onChange={handleCheckboxChange}>
                Exercise {id}
            </Checkbox>
            <div className="size mx-5 my-2 px-5">
                {question.split("\n").map((i, key) => {
                    return <div key={key}>{i}</div>;
                })}
                {!isSetLoading && <Upload
                    customRequest={customRequest}
                    onRemove={onRemove}
                    listType="picture"
                    defaultFileList={defaultFiles}
                >
                    <Button type='primary' disabled={completed} icon={<UploadOutlined />}>Upload File</Button>
                </Upload >}
            </div>
        </>
    );
}