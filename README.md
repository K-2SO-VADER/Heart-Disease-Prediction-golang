# Heart Disease (AI Class Project)

For AI class Project, this group will develop a predictive Heart Disease model based on UCI ML dataset You can download the dataset at: [https://archive.ics.uci.edu/ml/machine-learning-databases/heart-disease](https://archive.ics.uci.edu/ml/machine-learning-databases/heart-disease)

## Setting up the Environment

1. Programming Language: golang 1.7+
1. Visual Studio Code

### Setting up golang

Follow the instructions [here](https://golang.org/doc/install)

## Training and Prediction

## [golearn](https://github.com/sjwhitworth/golearn)

- A Machine Learning Library for golang based on [scikit Learn](http://scikit-learn.org/)

## How to run

```bash

go run *.go -a <algorithim> -p /path/to/dataset.csv

```

Where algorithm is either crossfold, decision_tree or knn

## Datasets

These are available in the datasets folder


### Decision Tree Results

```zsh

ID3 Performance (information gain)
Reference Class	True Positives	False Positives	True Negatives	Precision	Recall	F1 Score
---------------	--------------	---------------	--------------	---------	------	--------
three		0		0		142		NaN		0.0000	NaN
one		1		0		138		1.0000		0.0476	0.0909
zero		92		21		38		0.8142		0.9200	0.8638
two		9		36		112		0.2000		0.8182	0.3214
four		0		0		149		NaN		0.0000	NaN
Overall accuracy: 0.6415

ID3 Performance (information gain ratio)
Reference Class	True Positives	False Positives	True Negatives	Precision	Recall	F1 Score
---------------	--------------	---------------	--------------	---------	------	--------
three		5		11		131		0.3125		0.2941	0.3030
one		2		8		130		0.2000		0.0952	0.1290
zero		94		20		39		0.8246		0.9400	0.8785
two		3		16		132		0.1579		0.2727	0.2000
four		0		0		149		NaN		0.0000	NaN
Overall accuracy: 0.6541

ID3 Performance (gini index generator)
Reference Class	True Positives	False Positives	True Negatives	Precision	Recall	F1 Score
---------------	--------------	---------------	--------------	---------	------	--------
two		0		0		148		NaN		0.0000	NaN
four		0		0		149		NaN		0.0000	NaN
three		0		0		142		NaN		0.0000	NaN
one		14		32		106		0.3043		0.6667	0.4179
zero		94		19		40		0.8319		0.9400	0.8826
Overall accuracy: 0.6792

RandomTree Performance
Reference Class	True Positives	False Positives	True Negatives	Precision	Recall	F1 Score
---------------	--------------	---------------	--------------	---------	------	--------
three		7		8		134		0.4667		0.4118	0.4375
one		7		9		129		0.4375		0.3333	0.3784
zero		97		26		33		0.7886		0.9700	0.8700
two		0		0		148		NaN		0.0000	NaN
four		5		0		149		1.0000		0.5000	0.6667
Overall accuracy: 0.7296

RandomForest Performance
Reference Class	True Positives	False Positives	True Negatives	Precision	Recall	F1 Score
---------------	--------------	---------------	--------------	---------	------	--------
three		0		0		142		NaN		0.0000	NaN
one		0		0		138		NaN		0.0000	NaN
zero		100		59		0		0.6289		1.0000	0.7722
two		0		0		148		NaN		0.0000	NaN
four		0		0		149		NaN		0.0000	NaN
Overall accuracy: 0.6289
```

### KNN results

```zsh

Reference ClassnTrue Positives	False Positives	True Negatives	Precision	Recall	F1 Score
---------------	--------------	---------------	--------------	---------	------	--------
zero		98		56		1		0.6364		0.9800	0.7717
three		0		0		140		NaN		0.0000	NaN
one		0		0		135		NaN		0.0000	NaN
two		0		3		145		0.0000		0.0000	NaN
four		0		0		148		NaN		0.0000	NaN
Overall accuracy: 0.6242

```