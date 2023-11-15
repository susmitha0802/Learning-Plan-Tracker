import { ProgressBar } from 'react-bootstrap';
import { useUser } from '../../../contexts/UserContext';

export const Progress = ({ id, total, email }) => {
    const { getSubmittedCountFromLocalStorage } = useUser();
    const count = getSubmittedCountFromLocalStorage(email, id);
    const progress = Math.round(count / total * 100);

    return (
        <>
            <ProgressBar className="mb-4 size" now={progress} label={`${progress}%`} />
        </>
    )
}
