function main(input) {
	input.generated = 3;

	//while (true) {
	//input.generated++;
	//}
	for (let i = 0; i < 10000; ++i) {
		input.generated = input.generated + i;
	}

	return input;
}
