export const getExercises = (total, topics) => {
    topics.map(topic => {
        if (topic.exercises) {
            total = total + topic.exercises.length;
        }
        return total;
    })

    return total;
}