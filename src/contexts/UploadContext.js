import { createContext, useContext, useState } from "react";
import { CourseContext } from "../components/course/Course";
import axios from "axios";
import { message } from "antd";

const UploadContext = createContext();

export const useUpload = () => {
    return useContext(UploadContext);
}

export const UploadProvider = ({ children }) => {

    const [uploaded, setUploaded] = useState(false);
    const courseId = useContext(CourseContext);
    const handleRequest = ({ exerciseId }) => {
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
    }

    const handleRemove = () => {
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
    }


    const value = {
        uploaded,
        handleRequest,
        handleRemove,
        handleRemove.onRemove,
    }

    return (
        <UploadContext.Provider value={value}>
            {children}
        </UploadContext.Provider>
    );
}