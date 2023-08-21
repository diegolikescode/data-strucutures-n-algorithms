const AsyncQueue = require("./index");

const queue = new AsyncQueue();

async function asyncTask(duration) {
    setTimeout(function () {
        // why not?
        console.log(duration);
    }, duration);
}

function createInt32Array(length, min, max) {
    const newInt32Array = new Int32Array(length);
    for (let i = 0; i < length; i++) {
        newInt32Array[i] = Math.floor(Math.random() * (max - min + 1)) + min;
    }
    return newInt32Array;
}

// create an array with the task durations and call, then use asyncTask to create a bunch of tasks in my queue
const tasks = process.argv.slice(2).map((arg) => parseInt(arg));
const queueTasksDurations = createInt32Array(tasks[0], 3000, 6000);
const queueTasksDurations2 = createInt32Array(tasks[1], 8000, 28000);
const queueTasksDurations3 = createInt32Array(tasks[2], 30000, 62000);

for (
    let i = 0;
    i < queueTasksDurations.length &&
    i < queueTasksDurations2.length &&
    i < queueTasksDurations3.length;
    i++
) {
    if (queueTasksDurations[i])
        queue.enqueue(() => asyncTask(queueTasksDurations[i]));
    if (queueTasksDurations2[i])
        queue.enqueue(() => asyncTask(queueTasksDurations2[i]));
    if (queueTasksDurations3[i])
        queue.enqueue(() => asyncTask(queueTasksDurations3[i]));
}
