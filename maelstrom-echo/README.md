<https://fly.io/dist-sys/1/>
<https://github.com/jepsen-io/maelstrom/blob/main/demo/go>

```sh
brew install openjdk graphviz gnuplot

# Downlaod maelstrom
# https://github.com/jepsen-io/maelstrom/releases
# extract and move to ~/maelstrom

cd maelstrom-echo
export PATH=:$PATH:~/maelstrom
maelstrom test -w echo --bin ~/go/bin/maelstrom-echo --node-count 1 --time-limit 10
maelstrom serve
```
