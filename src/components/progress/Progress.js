import { ProgressBar } from 'react-bootstrap';
import { useUser } from '../../contexts/UserContext';
import { useAuth } from '../../contexts/AuthContext';

export const Progress = ({ id, total }) => {
    const { currentUser } = useAuth();
    const { getSubmittedCountFromLocalStorage } = useUser();

    const count = getSubmittedCountFromLocalStorage(currentUser.email, id);
    const progress = Math.round(count / total * 100);

    return (
        <>
            <ProgressBar now={progress} label={`${progress}%`} />
        </>
    )
}
