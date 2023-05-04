export function copy(text: string) {
	return new Promise<void>((resolve, reject) => {
		try {
			const input = document.createElement('input');
			input.style.position = 'absolute';
			input.style.top = '-9999px';
			input.value = text;
			document.body.appendChild(input);
			input.focus();
			input.select();
			input.setSelectionRange(0, input.value.length);
			setTimeout(() => {
				try {
					document.execCommand('copy');
					document.body.removeChild(input);
					resolve();
				} catch (err) {
					reject(err);
				}
			}, 0);
		} catch (err) {
			reject(err);
		}
	});
}
