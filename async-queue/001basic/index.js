class AsyncQueue {
    constructor() {
        this.queue = [];
        this.isProcessing = false;
        this.currId = 0;
    }

    enqueue(task) {
        this.currId += 1;
        task.id = this.currId;
        this.queue.push(task);
        if (!this.isProcessing) {
            this.process();
        }
    }

    async process() {
        if (this.queue.length === 0) {
            this.isProcessing = false;
            return;
        }

        this.isProcessing = true;
        const task = this.queue.shift();

        try {
            await task();
            // console.log(task.id);
        } catch (err) {
            console.log("error processin task:", err);
        }

        this.process();
    }
}

module.exports = AsyncQueue;
