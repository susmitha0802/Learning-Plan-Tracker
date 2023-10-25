import { ProgressBar } from 'react-bootstrap';
import { useCompleted } from '../../contexts/CompletedContext';

export const Progress = ({ name, total }) => {

    const { count } = useCompleted();
    const progress = Math.round(count / total * 100);

    return (
        <>
            <ProgressBar now={progress} label={`${progress}%`} />
        </>
    )
}
