#!/bin/bash
set -e

total_test_count=$(find /root/.duality/tests | grep .*\.sh$ --count)
test_count_start=$((ID * total_test_count / 4))
test_count_end=$(((ID + 1) * total_test_count / 4))

# wait for previous tests to finish (tests run in series)
bash /root/.duality/scripts/wait_for_tests_to_finish.sh $test_count_start

index=$((test_count_start + 1))
while [ $index -le $test_count_end ]
do
    echo "start test number: $index"

    # call test files by index number
    bash "$(find /root/.duality/tests | grep .*\.sh$ | head -$index | tail -1)"

    # note: attempting to record the processing time of any transaction in this way is difficult/impossible
    # as the logic for the transaction is not evoked immediately, it will be called and finished within
    # Tendermint's own process time on the current leading validator

    # this transaction send a memo of "completed-test-x" which all nodes may listen to to find test completion progress
    dualityd --home /testnet/dualityd tx bank send $(dualityd --home /testnet/dualityd keys show fred --output json | jq -r .address) $(dualityd --home /testnet/dualityd keys show fred --output json | jq -r .address) 1token -y --broadcast-mode block --output json --note "completed-test-$index"

    index=$((index + 1))
done