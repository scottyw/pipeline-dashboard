# Pipeline Dashboard

The pipeline dashboard is a tool for measuring the length of time that Jenkins builds with deep dependencies take to run, and then number of errors in those jobs.

# Running

Copy conf/config.example.toml to conf/config.toml and add your Prodcut and Jenkins Job information there.  A "product" is a codebase which may have multiple branches being monitored, with multiple Jenkins Pipelines
