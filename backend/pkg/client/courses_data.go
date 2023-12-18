package main

import (
	pb "lpt/pkg/proto"
)

var Courses = []*pb.CourseDetails{
	{
		Name:    "JavaScript",
		Caption: "Learn how to use JavaScript — a powerful and flexible programming language for adding website interactivity.",
		Logo:    "https://upload.wikimedia.org/wikipedia/commons/thumb/6/6a/JavaScript-logo.png/640px-JavaScript-logo.png",
		Time:    5,
	},
	{
		Name:    "React",
		Caption: "In this React course, you'll build powerful interactive applications with one of the most popular JavaScript libraries.",
		Logo:    "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/2300px-React-icon.svg.png",
		Time:    10,
	},
}

var Topics = []*pb.TopicDetails{
	{
		Name:     "Introduction",
		Resource: "https://www.codecademy.com/courses/introduction-to-javascript/informationals/learn-javascript-welcome",
		CourseId: 1,
	},
	{
		Name:     "Conditionals",
		Resource: "https://www.codecademy.com/courses/introduction-to-javascript/lessons/control-flow/exercisess/control-flow-intro",
		CourseId: 1,
	},
	{
		Name:     "Functions",
		Resource: "https://www.codecademy.com/courses/introduction-to-javascript/lessons/functions/exercisess/intro-to-functions",
		CourseId: 1,
	},
	{
		Name:     "Scope",
		Resource: "https://www.codecademy.com/courses/introduction-to-javascript/lessons/scope/exercisess/scope",
		CourseId: 1,
	},
	{
		Name:     "Arrays",
		Resource: "https://www.codecademy.com/courses/introduction-to-javascript/lessons/arrays/exercisess/arrays",
		CourseId: 1,
	},
	{
		Name:     "Loops",
		Resource: "https://www.codecademy.com/courses/introduction-to-javascript/lessons/loops/exercisess/loops",
		CourseId: 1,
	},
	{
		Name:     "Iterators",
		Resource: "https://www.codecademy.com/courses/introduction-to-javascript/lessons/higher-order-functions/exercisess/introduction",
		CourseId: 1,
	},
	{
		Name:     "Objects",
		Resource: "https://www.codecademy.com/courses/introduction-to-javascript/lessons/objects/exercisess/intro",
		CourseId: 1,
	},
	{
		Name:     "Classes",
		Resource: "https://www.javascripttutorial.net/javascript-class/",
		CourseId: 1,
	},
	{
		Name:     "Callbacks",
		Resource: "https://www.javascripttutorial.net/javascript-callback/",
		CourseId: 1,
	},
	{
		Name:     "Callback Queue and Event Loop",
		Resource: "https://frontendmasters.com/courses/javascript-hard-parts-v2/callback-queue-event-loop/",
		CourseId: 1,
	},
	{
		Name:     "Promises",
		Resource: "https://www.youtube.com/watch?v=YiYtwbnPfkY",
		CourseId: 1,
	},
	{
		Name:     "Demystifying Promises",
		Resource: "https://www.youtube.com/watch?v=cd3lPGfqMaw&t=1207s",
		CourseId: 1,
	},
	{
		Name:     "Async-Await",
		Resource: "https://app.pluralsight.com/id/signin/?redirectTo=https%3A%2F%2Fapp.pluralsight.com%2Flibrary%2Fcourses%2Fjavascript-promises-async-programming%2Ftable-of-contents",
		CourseId: 1,
	},
	{
		Name:     "NPM",
		Resource: "https://github.com/workshopper/how-to-npm",
		CourseId: 1,
	},
	{
		Name:     "Quick Start",
		Resource: "https://react.dev/learn",
		CourseId: 2,
	},
	{
		Name:     "Installation",
		Resource: "https://react.dev/learn/installation",
		CourseId: 2,
	},
	{
		Name:     "Describing the UI",
		Resource: "https://react.dev/learn/describing-the-ui",
		CourseId: 2,
	},
	{
		Name:     "Adding Interactivity",
		Resource: "https://react.dev/learn/adding-interactivity",
		CourseId: 2,
	},
	{
		Name:     "Managing State",
		Resource: "https://react.dev/learn/managing-state",
		CourseId: 2,
	},
	{
		Name:     "Escape Hatches",
		Resource: "https://react.dev/learn/escape-hatches",
		CourseId: 2,
	},
	{
		Name:     "Context",
		Resource: "https://react.dev/reference/react#context-hooks",
		CourseId: 2,
	},
	{
		Name:     "React Router",
		Resource: "https://www.youtube.com/playlist?list=PLC3y8-rFHvwjkxt8TOteFdT_YmzwpBlrG",
		CourseId: 2,
	},
	{
		Name:     "React Query",
		Resource: "https://www.youtube.com/playlist?list=PLC3y8-rFHvwjTELCrPrcZlo6blLBUspd2",
		CourseId: 2,
	},
	{
		Name:     "Recoil",
		Resource: "https://recoiljs.org/",
		CourseId: 2,
	},
}

var Exercises = []*pb.ExerciseDetails{
	{
		Question: "Write a condition that returns true for == but not === and explain why",
		TopicId:  2,
	},
	{
		Question: "Write a single if condition that returns true if x is greater than or equal to 10 and is even, or if x is less than 10 and is odd",
		TopicId:  2,
	},
	{
		Question: "Use a ternary operator to assign 'odd' or 'even' to a variable called 'output' by checking if a variable 'x' is odd or even",
		TopicId:  2,
	},
	{
		Question: "Using if/else if/else write a program that takes an input x (1-7) and prints corresponding weekday (1-Sunday etc). If input is invalid print invalid input.",
		TopicId:  2,
	},
	{
		Question: "Solve exercise 4 using switch case",
		TopicId:  2,
	},
	{
		Question: "What is difference between Function Declaration and Function Expression?",
		TopicId:  3,
	},
	{
		Question: "Write down the syntax rules of an Arrow function.",
		TopicId:  3,
	},
	{
		Question: "Using Function Declaration create a function which will add two numbers.",
		TopicId:  3,
	},
	{
		Question: "Using Function Expression create a function which will multiply two numbers.",
		TopicId:  3,
	},
	{
		Question: "Using Arrow Function create a function to subtract two numbers.",
		TopicId:  3,
	},
	{
		Question: "Create IIFE which will accept a string as argument and return the reversed string.",
		TopicId:  3,
	},
	{
		Question: "Guess the output - http://nicholasjohnson.com/javascript/javascript-for-programmers/exercisess/closure/",
		TopicId:  4,
	},
	{
		Question: "Is array passed by value or reference to a function. Justify your answer with an example.",
		TopicId:  5,
	},
	{
		Question: "Write a Javascript function to clone an array without using loops.",
		TopicId:  5,
	},
	{
		Question: "Breifly explain the shift, unshit, slice, splice, pop Array methods with examples",
		TopicId:  5,
	},
	{
		Question: "Using a for loop output the elements in reverse order [1,2,3,4,5]",
		TopicId:  6,
	},
	{
		Question: "What is the difference between for .. of and for .. in",
		TopicId:  6,
	},
	{
		Question: "Use map function to return the squares of the items in the below array [2, 4, 6, 8, 10].",
		TopicId:  7,
	},
	{
		Question: "Use filter function array to extract strings from the below array [1, 'two', 3, 'four', 5, 'Six', 7, 'Eight', 9, 'Ten'].",
		TopicId:  7,
	},
	{
		Question: "What are arguments of the reduce function. Use reduce function to return the sum of the items in the array [1, 3, 5, 7, 9].",
		TopicId:  7,
	},
	{
		Question: "Write a function which takes an object as input and print out its properties",
		TopicId:  8,
	},
	{
		Question: "How to delete property 'a' from object x={'a':1,'b':2}",
		TopicId:  8,
	},
	{
		Question: "Given x = {a:'1'}; y = x; does changing y.a change x.a? why? (looking for the answer - objects are pass by reference)",
		TopicId:  8,
	},
	{
		Question: "Copy x = {'a':1} to a variable y without creating a reference to x",
		TopicId:  8,
	},
	{
		Question: "Write a function to check if two objects are equal",
		TopicId:  8,
	},
	{
		Question: "Implement Queue using javascript class. It should have the enqueue, dequeue, front, isEmpty methods",
		TopicId:  9,
	},
	{
		Question: "Explain instanceOf operator with an example.",
		TopicId:  9,
	},
	{
		Question: "Solve https://www.hackerrank.com/contests/javascript-week2/challenges/js-callbacks",
		TopicId:  10,
	},
	{
		Question: "Briefly explain different Promise states.",
		TopicId:  12,
	},
	{
		Question: "Using promise to create a delayedLowerCase function which will transform a string to lower case after 5 seconds.\n Ex Input: \"BEAUTIFUL CODE\" Output: \"beautiful code\"",
		TopicId:  12,
	},
	{
		Question: "Briefly explain Promise Chaining. For the above created delayedLowerCase function chain a promise which will print the string and its length.\n Ex: Input: \"Beautiful Code\" Output: \"beautiful code - 14\"",
		TopicId:  12,
	},
	{
		Question: "Rewrite the below using Async-Await syntax\nfunction getProcessedData(url) {\n return downloadData(url) // returns a promise\n.catch(e => {\nreturn downloadFallbackData(url) // returns a promise\n})\n.then(v => {\nreturn processDataInWorker(v) // returns a promise\n})\n}",
		TopicId:  14,
	},
	{
		Question: "Create a component called Greetings which takes props and renders them on screen. Use this component in <App />\n - <Greetings /> should take a prop name and display Hello {prop}, Good Morning/Afternoon/Evening based on time.",
		TopicId:  18,
	},
	{
		Question: "Create component called ListView which Render a list of numbers\n - Default start and end numbers are 1 and 10, they should be able to be overridden from props\n - There shouldn't be any warnings in the browser console.\n - Explain the function and importance of key prop in a list",
		TopicId:  18,
	},
	{
		Question: "i. Build a basic signup form which takes name, email, password and confirm password fields.\nii. Add validations on submit button, validations are:\n- no field should be empty\n- password and confirm password should be same\n- email should have @ and . symbols in it.\niii. Show errors if the validations fail\nValues entered in the input fields should be stored in state.\nError values should not be stored in state but should be calculated from above state.\n",
		TopicId:  19,
	},
	{
		Question: "Explain when to lift state up to a parent component.",
		TopicId:  20,
	},
	{
		Question: "Lift the state from all above components to a component called Parent.\n- Call <Parent /> in App and all state should be managed in Parent\n- state updaters should be sent as props to children and should be updated from children.",
		TopicId:  20,
	},
	{
		Question: "Explain why functional components are better than class based components",
		TopicId:  21,
	},
	{
		Question: "Create a component called Counter that displays a number on screen with two buttons ( +, - ).\n- On clicking of those buttons, the number on screen should increase or decrease by 1.\n- Default value of that number should be 0 and parent should be able to override default value using as prop called defaultValue.",
		TopicId:  21,
	},
	{
		Question: "In the signup form you built above:\n- Have a switch out side the form component\n- If that switch is toggled input field for username should be shown. (Validations should be present for username)\n- If the switch is toggled again username shouldn't be shown. (Validations shouldn't be present for username)\nState for switch toggling should be in a context and shared using useContext between form and switch",
		TopicId:  22,
	},
	{
		Question: "Explain why context is useful even though we can lift state to parent and update it from child",
		TopicId:  22,
	},
	{
		Question: "Integrate react router to an app and have the following links. Each link should have its own component that occupies the full height and width of screen\ni. /home/:username - Should have the Greetings component from exercises 1 with default prop username from url,default value is \"User\"\nii. /login - Should have a login form from previous exercises\niii. /counter - Should have the counter component from previous exercises\n",
		TopicId:  23,
	},
	{
		Question: "Make a get call using useQuery to https://official-joke-api.appspot.com/random_joke and show jokes on screen.",
		TopicId:  24,
	},
	{
		Question: "Make a post call using useMutation from a free REST API service or mock API.",
		TopicId:  24,
	},
	{
		Question: "https://github.com/beautiful-code/ecommerce-react-exercises",
		TopicId:  25,
	},
}
