econstruct Project Description: Analyze the provided $PROJECT_DESCRIPTION$ to understand the core objectives, key features, functional and non-functional requirements, and the implied scope. Identify the major phases and components of the project.
2.  Generate Atomic Tasks: Based on your analysis, identify and create a list of atomic, distinct, and highly specific tasks. An atomic task is a single, assignable unit of work that cannot be reasonably broken down further into subtasks. Each task should be a direct, practical step a developer would take. For example, instead of a task like "Implement Backend," create separate tasks like "Implement User Authentication," "Develop CRUD API for Products," and "Integrate Payment Gateway."
3.  Title, Description, and Acceptance Criteria: For each atomic task, create a concise and descriptive title. Then, write a detailed description that clearly outlines what needs to be achieved. Crucially, include a bulleted list of acceptance criteria at the end of each description. These criteria should be specific, verifiable conditions that must be met for the task to be considered complete. They define the "definition of done" for the task.
4.  Assign Priority Level: For each task, assign a priority level: `high`, `medium`, or `low`.
5.  Estimate Expected Time: For each task, provide an `expected_time_hours` estimate in numerical hours. This estimate should be a realistic projection for a single atomic task.
6.  Format as JSON: Structure the entire output as a single JSON array, where each element represents a distinct task
.
Output format: 
The output must be only a valid JSON array of objects. No extra explanations before or after the JSON array of objects. Each object in the array will represent a distinct task and must adhere to the following schema:

JSON

[
  {
    "task_title": "String: A concise, actionable title for the task (e.g., 'Implement CRUD API for Users').",
    "task_description": "String: A detailed explanation of the task, including what needs to be done and its scope. Must end with a bulleted list of acceptance criteria defining the 'definition of done'.",
    "priority": "String: The priority level of the task ('high', 'medium', or 'low').",
    "expected_time_hours": "Number: Estimated time in hours (integer or float, e.g., 8, 16.5)."
  },
  // ... more task objects
]


$PROJECT_DESCRIPTION$ - 