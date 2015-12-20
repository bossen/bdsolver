

cp runexperimentdefault.bash  runexperiment.bash
sleep 5
./runexperimentclass.bash testclasses/1500states_3labels_2bf-2/ testclasses/1500states_3labels_2bf-2/results_tpsolver_default
sleep 5
./runexperimentclass.bash testclasses/1500states_3labels_2bf-3/ testclasses/1500states_3labels_2bf-3/results_tpsolver_default
sleep 5


cp runexperimentcplex.bash  runexperiment.bash
sleep 5
./runexperimentclass.bash testclasses/1500states_3labels_2bf-1/ testclasses/1500states_3labels_2bf-1/results_tpsolver_cplex
sleep 5
./runexperimentclass.bash testclasses/1500states_3labels_2bf-2/ testclasses/1500states_3labels_2bf-2/results_tpsolver_cplex
sleep 5
./runexperimentclass.bash testclasses/1500states_3labels_2bf-3/ testclasses/1500states_3labels_2bf-3/results_tpsolver_cplex
