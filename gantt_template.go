package gdag

var ganttTemplate = `<html>
    <body>
        <script src="https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js"></script>
        <script>
            mermaid.initialize({ startOnLoad: true });
        </script>

        Here is one mermaid diagram:
        <div class="mermaid">
            gantt
            dateFormat  YYYY-MM-DD

            %s
        </div>
    </body>
</html>`
