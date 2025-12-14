#DAYS = 0 1 2 3
DAYS = $(subst cmd/day,,$(wildcard cmd/day*))
TARGETS = $(subst cmd/,out/,$(wildcard cmd/day*))

# enable for a short test
ifeq 'x' 'y'
test:
	echo $(TARGETS)
	echo $(DAYS)
endif

all: $(TARGETS) out out/util out/optimize
	echo All days implemented: $(DAYS)
	echo All targets: $(TARGETS)

.PHONY : clean
clean:
	rm -rf out/

out:
	mkdir out

out/util: cmd/util/*
	go test adventofcode/year2025/cmd/util
	touch out/util

out/optimize: cmd/optimize/*
	go test adventofcode/year2025/cmd/optimize
	touch out/optimize

# use variable $(1) day number
define MAKEDAY
out/day$(1): cmd/day$(1)/day$(1).go cmd/day$(1)/*
	echo "### Testing Day $(1) ###"
	go test adventofcode/year2025/cmd/day$(1)
	echo "### Running Day $(1) ###"
	go run $$< > $$@ 2> $$@.log
	cat $$@
endef

$(foreach DAY,$(DAYS),$(eval $(call MAKEDAY,$(DAY))))
