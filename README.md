# port-service

# Golang microservices assignment
Read the instructions carefully.
Your ability to understand and follow instructions is part of the evaluation.
Your project will be evaluated from various angles code structure, design, API signature, testing, documentation. Importance should be given to the structure of the project and design of the packages.
The correctness of algorithms comes second in relation to API design and idiomatic golang. Your code needs to be eas!
to follow and easy to maintain for other developers.
You should focus on writing your best code,
don't write a lot of code. It doesn't matter if the service is
incomplete, we will be able to evaluate what matters to XXX from a few packages interacting with each other.
Code reviewers might engage with your repository via github issues (or equivalent), this is entirely part of the interview process, so try to answer their questions and provide possible code changes to fix issues.

## Time limits
The evaluation result of the test is not linked to how much time you spend on it.
Please DO NOT spend more than 2 or 3 hours doing it, if you haven't complete the task simply submit as is.
Successful applications show us that 2 or 3 hours are more than enough to cover all the evaluation points below. Prefer writing correct code over writing lots of code. The test should be the showcase of your application design ability and code quality,
there are no extra points for repetition, so if you think you
already wrote enough code to demonstrate what you will do for a specific pattern, use comments in similar code blocks.
This assignment is meant to evaluate the golang proficiency of full-time engineers.
Your code structure should follow microservices best practices and our evaluation will focus primarily on your ability to follow good design principles and les: on correctness and completeness of algorithms. During the face to face interview you will have the opportunity to explain your design choices and provide justifications for the parts that you omitted.
**Do not mention XXXXX or XXXXXX anywhere on the code or repository name.**

## Evaluation points in order of importance
- use of clean code, which is self documenting
- use of packages to achieve separation of concerns
- use of domain driven design
- use of golang idiomatic principles
- use of docker
- tests for business logic
- use of code quality checks such as linters and build tools
- use of git with appropriate commit messages
- documentation: README and inline code comments
- you MUST use go modules and a version of go >= 1.19
Results: please share a git repository with us containing your implementation. We prefer Github as we are familiar with the issues PR workflow.
Level of experience targeted: EXPERT
Avoid using frameworks such as go-kit, go-micro, gin, since they help infrastructure and application design, which is one of the evaluation points.
You can use libraries for testing, mapping, validation, etc.
If you have questions about the test, please draw your own conclusions according to what make sense from your point of you.

Good luck.
Time limitations: there are no hard time limits. Once again, do not spend more than 2 or 3 hours.

## Technical test
- Given a file with ports data (ports. json), write a port domain service that either creates a new record in a database, or updates the existing one (Hint: no need for delete or other methods).
- The file is of unknown size, it can contain several millions of records, you will not be able to read the entire file at once.
- The service has limited resources available (e.g. 200MB ram).
- The end result should be a storage containing the ports, representing the latest version found in the JSON. (Hint: use an in memory database to save time and avoid complexity).
- A Dockerfile should be used to contain and run the service (Hint: extra points for avoiding code building in docker).
- Provide at least one example per test type that you think are needed for your assignment. This will allow the reviewer to evaluate your critical thinking as well as your knowledge about testing.
- Your readme.md should explain how to run your program and test it.
- The service should handle certain signals correctly (e.g. a TERM or KILL signal should result in a graceful shutdown).
